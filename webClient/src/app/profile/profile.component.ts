import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router, ActivatedRoute, RouterStateSnapshot }       from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import {ProfileService,Profile} from '../profile.service'

@Component({
  moduleId: module.id,
  selector: 'app-profile',
  templateUrl: 'profile.component.html',
  styleUrls: ['profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

  private paramSub: Subscription
  private profile: Profile
  private counties = ["Waterford"]
  private profileAreas = {"Waterford":["Tramore Waterford City West","Waterford City East","Waterford City South"]}
  private updated = false
  private newSignUp = false

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private service: ProfileService){}

  ngOnInit() {
     console.log(this.router.routerState.queryParams);
     this.router.routerState.queryParams.subscribe(params => {
       this.newSignUp = params["newSignUp"] === "true";
     });
     this.paramSub = this.route.params.subscribe(params => {
       let id = params['id'];
       
       this.service.getProfile(id).then((u)=>{
         if(! u){
           console.error("user not defined", u);
           return;
         }
         this.profile = new Profile(u.id,u.county,u.area,u.email,u.firstName,u.secondName);
       }).catch((e)=>console.error);
     });
  }
  ngOnDestroy(){
    this.paramSub.unsubscribe();
  }

  profileSubmit(){
    this.service.updateProfile(this.profile)
    .then((p)=>{
      console.log("updated profile", p)
      this.updated = true; 
      setTimeout(()=>{
        this.updated = false;
      },750)
    })
    .catch((e)=>console.error)
  }
}

