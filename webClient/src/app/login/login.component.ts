import { Component, OnInit, NgZone } from '@angular/core';
import {LoginService, Login} from './login.service'
import {ProfileService,UserData} from '../profile.service'
import {ROUTER_DIRECTIVES, Router} from '@angular/router'

declare var gapi:any;
@Component({
  moduleId: module.id,
  selector: 'app-login',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.css'],
  providers:[LoginService]
})
export class LoginComponent implements OnInit {

  constructor(private _zone: NgZone,private loginService: LoginService, private profileService: ProfileService, private router: Router ) { }
  userAuthToken = null;
  userDisplayName = "empty";
  userEmail = "";


  ngAfterViewInit() {
    // Converts the Google login button stub to an actual button.
    var self = this;
    gapi.load('auth2', function() {
      gapi.auth2.init();
      gapi.signin2.render(
          "google-signin",
          {
            "onSuccess": self.onGoogleLoginSuccess,
            "scope": "profile",
            "theme": "dark"
          });
    });
  }

  // Triggered after a user successfully logs in using the Google external
  // login provider.
  onGoogleLoginSuccess = (loggedInUser) => {
    this._zone.run(() => {
      this.userAuthToken = loggedInUser.getAuthResponse().id_token;
      this.userDisplayName = loggedInUser.getBasicProfile().getName();
      this.userEmail = loggedInUser.getBasicProfile().getEmail();
      let login = new Login(this.userAuthToken,this.userEmail,"google");
      this.loginService.login(login)
      .then(this.handleLoginSuccess(this.profileService, this.router))
      .catch((err)=>{
        console.log("error logging in ", err);
      });
    });
  }

  handleLoginSuccess(profileService, router){
    return function (res){
        console.log("loginSuccess",res);
        let ud = new UserData(res.id,res.token,res.type);
        let pStore = profileService.storeUserData(ud)
        pStore.then(()=>{
          router.navigate(["/profile/"+ud.id]);
        });
        pStore.catch((err)=>console.error);
    }
  }

  login(){
    console.log("login called", this.userEmail, this.userAuthToken);
    let login = new Login(this.userAuthToken,this.userEmail,"local");
      this.loginService.login(login)
      .then(this.handleLoginSuccess(this.profileService, this.router))
      .catch((err)=>{
        console.log("error logging in ", err);
      });
  }

  ngOnInit() {
    
  }

}
