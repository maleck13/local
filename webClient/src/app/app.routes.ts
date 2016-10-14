import { provideRouter, RouterConfig } from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {SignupComponent} from "./signup/signup.component";
import {ProfileComponent} from "./profile/profile.component";
import {AdminCoucillorComponent} from "./admin/admin.councillor.component";
import {CouncillorsComponent} from "./councillors/councillors.component";
const routes: RouterConfig = [
    { path: '', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'signup', component: SignupComponent },
    {path: 'profile/:id', component: ProfileComponent},
    {path: 'admin/councillor', component:AdminCoucillorComponent},
    {path: 'councillors',component:CouncillorsComponent}
];

export const appRouterProviders = [
    provideRouter(routes)
];
