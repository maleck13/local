import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import {Councillor, CouncillorsService, CouncillorCommunication} from '../councillors/councillors.service'
import {ActivatedRoute, RouterStateSnapshot, Router }       from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import {ProfileService,UserData} from '../profile.service'
import {Response} from '@angular/http'
import {CommunicationsService} from '../communications.service'

@Component({
  selector: 'councillor',
  templateUrl: 'councillor.component.html',
  styleUrls: ['councillor.component.css'],
  providers: [CouncillorsService, CommunicationsService]
})
export class CouncillorComponent implements OnInit, OnDestroy {

  constructor(private route: ActivatedRoute, private service: CouncillorsService
    , private profile: ProfileService, private router: Router,
    private communications: CommunicationsService) { }

  @Input() councillors: Councillor[]

  private councillor: Councillor;
  private paramSub: Subscription;
  private communication: CouncillorCommunication;
  private selectedCommunication: CouncillorCommunication;
  private subComms: CouncillorCommunication[];
  private councillorComunications: CouncillorCommunication[];
  private error: Response;
  private user: UserData;

  ngOnInit() {
    this.paramSub = this.route.params.subscribe(params => {
      let id = params["id"];
      this.profile.getTokenHeader()
        .then((auth) => {
          this.service.councillor(id, auth)
            .then((c) => {
              this.councillor = c;
              this.communications.listForUser(c.id,auth,null)
              .then((comms)=>{
                this.councillorComunications = comms;
              })
              .catch((err)=>{
                this.handleErrorResponse(err);  
              })
            })
            .catch((err) => {
              this.handleErrorResponse(err);
            });
        });
    });
     this.profile.getUserData()
     .then((u)=>this.user = u);
  }
  ngOnDestroy() {
    this.paramSub.unsubscribe();
  }

  communicate(c: Councillor) {

    this.communication = new CouncillorCommunication("","", c.id, "", true,"email",null);
  }
  cancelCommunication(ev: any) {
    ev.preventDefault()
    this.communication = null;

  }

  openCommunication(councillorID:string,commID: string, index:number){
    console.log("openCommunication", commID);
    //call out to api to list comms by commID
    this.profile.getTokenHeader()
        .then((auth) => {
          return this.communications.listForUser(councillorID,auth,commID)
        })
        .then((comms)=>{
          this.selectedCommunication = this.councillorComunications[index];
          this.subComms = comms.filter((c)=>{
            if (c.id != this.selectedCommunication.id){
              return true
            }
          });
        })
        .catch((e)=>console.error);
    
  }

  sendCommunication(councillor: Councillor) {
    console.log("c", councillor, this.communication);
    this.profile.getTokenHeader()
      .then((auth) => {
        this.communications.communicate(this.communication, auth)
          .then((res) => {
            console.log("sent communication");
          })
          .catch((err) => {
            this.handleErrorResponse(err);
          });
      });
  }
  handleErrorResponse(err: Response) {
    this.error = err;
    setTimeout(() => {
      if (err && err.status === 401) {
        this.profile.deleteUserData().
          then(() => {
            this.router.navigate(["/login"]);
          });
      }
    }, 1500);
  }



}
