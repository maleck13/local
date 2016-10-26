import { Directive , ElementRef,Renderer} from '@angular/core';
import {ProfileService} from './profile.service'

@Directive({
  selector: '[appIscouncillor]'
})
export class IscouncillorDirective {

  constructor(el: ElementRef, renderer: Renderer, private profileSevice:ProfileService) {
     console.log("is logged in Directive");
     profileSevice.getUserData().then((data)=>{
     })
  }

}
