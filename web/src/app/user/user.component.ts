import { Component, OnInit } from '@angular/core';
import {UserService} from '../user.service'
import {LoginService} from '../login/login.service'
import { Router, ActivatedRoute, RouterStateSnapshot }       from '@angular/router';
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css'],
  providers:[UserService]
})
export class UserComponent implements OnInit {

  constructor(private userService: UserService,
  private route: ActivatedRoute,
    private router: Router,
    private loginService: LoginService) { }

    private key:String
    private uid:String 
    private newpassword:String
    private error:String
    private success:String 

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.key = params["key"];
      this.uid = params["uid"];
    });
  }

  resetPassword(){
    console.log("reset password ", this.newpassword);
    this.userService.resetPassword(this.key,this.newpassword)
    .then((res)=>{
        this.success = "Password Successfully updated. Please log in."
    })
    .catch((err)=>{
      this.error = err
    })
  }



}
