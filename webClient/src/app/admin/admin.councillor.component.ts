import { Component, OnInit, NgZone, EventEmitter } from '@angular/core';
import {Councillor, AdminCouncillorService} from './admin.councillor.service'
import { UPLOAD_DIRECTIVES } from 'ng2-uploader';
import {ProfileService} from '../profile.service'

@Component({
  moduleId: module.id,
  selector: 'app-admin',
  templateUrl: 'admin.councillor.component.html',
  styleUrls: ['admin.councillor.component.css'],
  providers: [AdminCouncillorService,ProfileService],
  directives: [UPLOAD_DIRECTIVES]
})
export class AdminCoucillorComponent implements OnInit {

  constructor(private service: AdminCouncillorService, private profile: ProfileService) { }

  private councillor: Councillor = new Councillor("", "", "","","","","","","");
  private areas = ["Tramore Waterford City West", "Waterford City East", "Waterford City South"];
  private zone: NgZone;
  private basicOptions: Object;
  private progress: number = 0;
  private response: any = {};
  private uploadEvents: EventEmitter<any> = new EventEmitter();



  handleUpload(data: any): void {
    console.log(data);
    this.zone.runGuarded(() => {
      this.response = data;
      this.progress = Math.floor(data.progress.percent / 100);
    });
    this.zone.onError.subscribe((err)=>{
      console.log("zone error",err);
    })
  }

  ngOnInit() {
    this.zone = new NgZone({ enableLongStackTrace: false });
    this.basicOptions = {
      url: '/admin/councillor',
      calculateSpeed: true,
      autoUpload: false,
      previewUrl: false,
      data: this.councillor,
      authToken: '',//added after we get the token
      authTokenPrefix: 'Bearer'
    };
    this.profile.getJwtToken().then((token)=>{
      this.basicOptions.authToken = token;
    })
    .catch((err)=>console.error);
  }

  add() {
    this.uploadEvents.emit('startUpload');
    this.uploadEvents.subscribe((data)=>{
      console.log(data);
    }, (err)=>{
      console.log("upload error ", err);
    })
  }

}
