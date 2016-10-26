var router_1 = require('@angular/router');
var login_component_1 = require("./login/login.component");
var home_component_1 = require("./home/home.component");
var signup_component_1 = require("./signup/signup.component");
var routes = [
    { path: '', component: home_component_1.HomeComponent },
    { path: 'login', component: login_component_1.LoginComponent },
    { path: 'signup', component: signup_component_1.SignupComponent }
];
exports.appRouterProviders = [
    router_1.provideRouter(routes)
];
//# sourceMappingURL=app.routes.js.map