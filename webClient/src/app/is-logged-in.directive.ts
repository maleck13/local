import { Directive, ElementRef, Renderer } from '@angular/core';
import {ProfileService} from './profile.service'

@Directive({
  selector: '[isLoggedIn]'
})
export class IsLoggedIn {

   constructor(el: ElementRef, renderer: Renderer, private profileSevice:ProfileService) {
     console.log("is logged in Directive");
     profileSevice.getUserData().then((data)=>{
       if(data && data.id){
        renderer.setElementStyle(el.nativeElement, 'display', 'none');  
       }else{
         renderer.setElementStyle(el.nativeElement, 'display', 'block');
       }
     }).catch((err)=>console.error);
    }

}
