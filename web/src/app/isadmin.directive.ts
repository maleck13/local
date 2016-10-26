import { Directive, ElementRef, Renderer } from '@angular/core';
import {ProfileService} from './profile.service'

@Directive({
  selector: '[isadmin]'
})
export class Isadmin {

    constructor(el: ElementRef, renderer: Renderer, private profileSevice:ProfileService) {
     
     profileSevice.getUserData().then((data)=>{
       console.log("isadmin ", data);
       renderer.setElementStyle(el.nativeElement, 'display', 'none');  
       if (data && data.type == "admin"){
         renderer.setElementStyle(el.nativeElement, 'display', 'block');
       }
     }).catch((err)=>console.error);
    }

}