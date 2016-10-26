import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import {Councillor, CouncillorsService, CouncillorCommunication} from '../councillors/councillors.service'
import {ActivatedRoute, RouterStateSnapshot, Router }       from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import {ProfileService} from '../profile.service'
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

  private councillor: Councillor
  private paramSub: Subscription
  private communication: CouncillorCommunication
  private councillorComunications: CouncillorCommunication[]
  private error: Response

  ngOnInit() {
    this.paramSub = this.route.params.subscribe(params => {
      let id = params["id"];
      this.profile.getTokenHeader()
        .then((auth) => {
          this.service.councillor(id, auth)
            .then((c) => {
              this.councillor = c;
              this.communications.listForUser(c.id,auth)
              .then((comms)=>{
                console.log(comms)
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
  }
  ngOnDestroy() {
    this.paramSub.unsubscribe();
  }

  communicate(c: Councillor) {

    this.communication = new CouncillorCommunication("", c.id, "", true,"email");
  }
  cancelCommunication(ev: any) {
    ev.preventDefault()
    this.communication = null;

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
