import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router, ActivatedRoute, RouterStateSnapshot }       from '@angular/router';
import { Subscription } from 'rxjs/Subscription';
import {ProfileService, Profile, UserData} from '../profile.service'
import {Response} from '@angular/http'

@Component({
  selector: 'app-profile',
  templateUrl: 'profile.component.html',
  styleUrls: ['profile.component.css']
})
export class ProfileComponent implements OnInit, OnDestroy {

  private paramSub: Subscription
  private profile: Profile
  private counties = ["Waterford"]
  private profileAreas = { "Waterford": ["Tramore Waterford City West", "Waterford City East", "Waterford City South"] }
  private updated = false
  private newSignUp = false
  private error: Response

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private service: ProfileService) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.newSignUp = params["newSignUp"] === "true";
    });
    this.paramSub = this.route.params.subscribe(params => {
      let id = params['id'];

      this.service.getProfile(id).then((u) => {
        if (!u) {
          console.error("user not defined", u);
          return;
        }
        this.profile = new Profile(u.id, u.county, u.area, u.email, u.firstName, u.secondName);
      }).catch((e) => {
        this.handleErrorResponse(e);
      });
    });
  }
  ngOnDestroy() {
    this.paramSub.unsubscribe();
  }

  handleErrorResponse(err: Response) {
    this.error = err;
    setTimeout(() => {
      if (err && err.status === 401) {
        this.service.deleteUserData().
          then(() => {
            this.router.navigate(["/login"]);
          });
      }
    }, 1500);
  }

  profileSubmit() {
    this.service.updateProfile(this.profile)
      .then((p) => {
        this.service.getUserData()
          .then((ud) => {
            ud.county = this.profile.county || "";
            ud.area = this.profile.area || "";
            return ud;
          })
          .then((ud) => {
            this.service.storeUserData(ud)
              .catch((err) => {
                this.handleErrorResponse(err);
              });
          })
      })
      .then((p) => {
        this.updated = true;
        setTimeout(() => {
          this.updated = false;
        }, 750)
      })
      .catch((e) => {
        this.handleErrorResponse(e);
      })
  }
}

