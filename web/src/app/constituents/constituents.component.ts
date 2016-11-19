import { Component, OnInit } from '@angular/core';
import {ConstituentsService,Constituent} from '../constituents.service'
import {ProfileService} from '../profile.service'

@Component({
  selector: 'app-constituents',
  templateUrl: './constituents.component.html',
  styleUrls: ['./constituents.component.css'],
  providers:[ConstituentsService]

})

export class ConstituentsComponent implements OnInit {

  constructor(
    private cs: ConstituentsService,
   private profileService: ProfileService) { }

   private constituents: Constituent[]


   selectConstituent(){
     console.log("selected");
   }

  ngOnInit() {
    console.log("ngOnInit")
    let auth = null;
    this.profileService.getTokenHeader()
    .then((a)=>{
          auth = a;
          return this.profileService.getUserData();
    })
    .then((ud)=>{
      return this.cs.constituents(ud.id,auth)
    })
    .then((data)=>{
      console.log(data);
      this.constituents = data;
    })
    .catch((e)=>console.error);
  }

}
