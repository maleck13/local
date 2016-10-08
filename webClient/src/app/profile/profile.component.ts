import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router, ActivatedRoute }       from '@angular/router';
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
  private profileAreas = ["Tramore Waterford City West","Waterford City East","Waterford City South"]
  private updated = false

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private service: ProfileService){}

  ngOnInit() {
     this.paramSub = this.route.params.subscribe(params => {
       let id = params['id']; 
       this.service.getProfile(id).then((u)=>{
         if(! u){
           console.log("user not defined", u);
           return;
         }
         console.log("user ", u);
         this.profile = new Profile(u.id,u.area,u.email,u.firstName,u.secondName);
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

