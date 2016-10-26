import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router, ActivatedRoute, RouterStateSnapshot }       from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import {ProfileService} from '../profile.service'
import {CouncillorsService, Councillor} from './councillors.service'
import {CouncillorComponent} from '../councillor/councillor.component'
import {Response} from '@angular/http'

@Component({
  selector: 'app-councillors',
  templateUrl: 'councillors.component.html',
  styleUrls: ['councillors.component.css'],
  providers: [CouncillorsService]
})
export class CouncillorsComponent implements OnInit , OnDestroy{

  constructor(private route: ActivatedRoute,
    private router: Router,
    private profile: ProfileService,
    private service: CouncillorsService) { }

  public councillors: Councillor[]
  private paramSub: Subscription
  private error: Response
  // init component when angular is ready
  ngOnInit() {
    this.paramSub = this.route.params.subscribe(params => {
      let county = params['county'];
      let area = params['area'];
      this.profile.getTokenHeader()
        .then((auth) => {
          this.service.councillors(county, area, auth)
            .then((res) => {
              this.councillors = res;
            })
            .catch((err) => {
              this.handleErrorResponse(err);
            });
        })

    });
  }
  // unsubscribe from observable
  ngOnDestroy(){
    this.paramSub.unsubscribe();
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

  selectCouncillor(county, id: string) {
    this.router.navigate(["/councillor/" + county + "/" + id])
  }

}
