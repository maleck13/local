import { Component, OnInit, Input, OnDestroy, NgZone, EventEmitter } from '@angular/core';
import { Councillor } from '../../admin/admin.councillor.service'
import { CouncillorsService } from '../../councillors/councillors.service'
import { Subscription } from 'rxjs/Subscription';
import { Router, ActivatedRoute, RouterStateSnapshot } from '@angular/router';
import { ProfileService, Profile, UserData } from '../../profile.service'

@Component({
  selector: 'councillor',
  templateUrl: 'councillor.profile.component.html',
  styleUrls: ['councillor.profile.component.css'],
  providers: [CouncillorsService]
})

export class CouncillorProfileComponent implements OnInit {
  private paramSub: Subscription

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private service: CouncillorsService,
    private profileService: ProfileService) { }

  private councillor: Councillor
  private zone: NgZone;
  private basicOptions: any = {};
  private progress: number = 0;
  private response: any = {};
  private uploadEvents: EventEmitter<any> = new EventEmitter();
  private areas = ["Tramore Waterford City West", "Waterford City East", "Waterford City South"];
  private counties = ["Waterford"];

  handleUpload(data: any): void {
    this.zone.runGuarded(() => {
      try{
      let res = JSON.parse(data.response);
      this.councillor.image = res.image;
      }catch(e){
        console.error(e);
      }
      this.response = data;
      this.progress = Math.floor(data.progress.percent / 100);
    });
    this.zone.onError.subscribe((err) => {
      console.log("zone error", err);
    })
  }

  ngOnInit() {
    this.paramSub = this.route.params.subscribe(params => {
      let id = params['id'];
      this.profileService.getTokenHeader()
        .then((auth) => {
          return this.service.councillor(id, auth);
        })
        .then((councillor) => {
          this.councillor = councillor;
          this.zone = new NgZone({ enableLongStackTrace: false });
          this.profileService.getJwtToken()
            .then((token) => {
              this.basicOptions = {
                url: '/councillors/' + id + '/image',
                calculateSpeed: true,
                autoUpload: true,
                previewUrl: false,
                authToken: token,//added after we get the token
                authTokenPrefix: 'Bearer'
              };
            });
        })
        .catch((e) => console.error);
    });
  }
  
  update(){
    this.profileService.getTokenHeader()
    .then((auth)=>{
      if (! auth){
        console.log("no auth from local storage");
      }
      console.log(auth);
      return this.service.update(this.councillor,auth);
    })
    .catch((err)=>console.error);
  }

}