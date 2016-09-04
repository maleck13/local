import { Directive, ElementRef, Renderer } from '@angular/core';
import {ProfileService} from './profile.service'

@Directive({
  selector: '[isguest]'
})
export class Isguest {

    constructor(el: ElementRef, renderer: Renderer, private profileSevice:ProfileService) {
     
     profileSevice.getUserData().then((data)=>{
       console.log("is guest Directive", data);
       if(! data ){
        renderer.setElementStyle(el.nativeElement, 'display', 'none');  
       }else{
         renderer.setElementStyle(el.nativeElement, 'display', 'block');
       }
     }).catch((err)=>console.error);
    }

}