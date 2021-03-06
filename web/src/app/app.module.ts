import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { UPLOAD_DIRECTIVES } from 'ng2-uploader';
import { FormsModule }    from '@angular/forms';
import { RouterModule }   from '@angular/router';
import { AppComponent }  from './app.component';
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {SignupComponent} from "./signup/signup.component";
import {ProfileComponent} from "./profile/profile.component";
import {AdminCoucillorComponent} from "./admin/admin.councillor.component";
import {CouncillorsComponent} from "./councillors/councillors.component";
import {CouncillorComponent} from "./councillor/councillor.component";
import {IsLoggedIn} from './is-logged-in.directive';
import {Isguest} from './isguest.directive';
import {Isadmin} from './Isadmin.directive';
import {HttpModule}from '@angular/http';
import { UserComponent } from './user/user.component';
import { IscouncillorDirective } from './iscouncillor.directive';
import { ConstituentsComponent } from './constituents/constituents.component';
import { CouncillorProfileComponent } from './councillor/profile/councillor.profile.component';

@NgModule({
  imports: [ 
      BrowserModule,
      FormsModule ,
      HttpModule,
      RouterModule.forRoot([
        { path: '', component: HomeComponent },
        { path: 'login', component: LoginComponent },
        { path: 'signup', component: SignupComponent },
        {path: 'local/profile/:id', component: ProfileComponent},
        {path: 'councillor/profile/:id', component: CouncillorProfileComponent},
        {path: 'admin/councillor', component:AdminCoucillorComponent}, //TODO CHANGE THIS TO USE THE CouncillorProfileComponent
        {path: 'councillors/:county/:area',component:CouncillorsComponent},
        {path: 'councillor/:county/:id',component:CouncillorComponent},
        {path: 'constituents/:county/:area',component:ConstituentsComponent},
        {path:'passwordreset', component:UserComponent}
        ])],
  declarations: [ 
      AppComponent,
      HomeComponent,
      LoginComponent,
      SignupComponent,
      ProfileComponent,
      AdminCoucillorComponent,
      CouncillorsComponent,
      CouncillorComponent,
      CouncillorProfileComponent,
      UPLOAD_DIRECTIVES ,
      IsLoggedIn,
      Isguest,
      Isadmin,
      UserComponent,
      IscouncillorDirective,
      ConstituentsComponent],
  bootstrap: [ AppComponent ]
})
export class AppModule { }