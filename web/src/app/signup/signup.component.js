var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var core_1 = require('@angular/core');
var signup_service_1 = require("./signup.service");
var signup_service_2 = require("./signup.service");
var SignupComponent = (function () {
    function SignupComponent(_zone, signUpService) {
        var _this = this;
        this._zone = _zone;
        this.signUpService = signUpService;
        // Triggered after a user successfully logs in using the Google external
        // login provider.
        this.onGoogleLoginSuccess = function (loggedInUser) {
            console.log("here");
            _this._zone.run(function () {
                _this.userAuthToken = loggedInUser.getAuthResponse().id_token;
                _this.userDisplayName = loggedInUser.getBasicProfile().getName();
                _this.userEmail = loggedInUser.getBasicProfile().getEmail();
                _this.givenName = loggedInUser.getBasicProfile().getGivenName();
                _this.familyName = loggedInUser.getBasicProfile().getFamilyName();
                console.log("here", loggedInUser.getBasicProfile());
                console.log(_this.userDisplayName, _this.userAuthToken, _this.userEmail);
                var model = new signup_service_2.Signup("google", _this.userAuthToken, _this.userEmail, _this.givenName, _this.familyName);
                _this.signUpService.signUp(model)
                    .then(function (res) {
                    console.log(res);
                    //redirect to complete sign up
                })
                    .catch(function (err) {
                    console.log("sign up error", err);
                    if (err.status && err.status == 409) {
                        _this.message = "You have already registered. Please sign in";
                    }
                    else {
                        _this.message = "unexpected error occurred. Please try again later";
                    }
                });
            });
        };
    }
    SignupComponent.prototype.ngOnInit = function () {
        this.message = "Signup Message";
    };
    SignupComponent.prototype.ngAfterViewInit = function () {
        // Converts the Google login button stub to an actual button.
        var self = this;
        gapi.load('auth2', function () {
            gapi.auth2.init();
            gapi.signin2.render("google-signin", {
                "onSuccess": self.onGoogleLoginSuccess,
                "scope": "profile",
                "theme": "dark"
            });
        });
    };
    SignupComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'app-signup',
            templateUrl: 'signup.component.html',
            styleUrls: ['signup.component.css']
        }), 
        __metadata('design:paramtypes', [core_1.NgZone, signup_service_1.SignupService])
    ], SignupComponent);
    return SignupComponent;
})();
exports.SignupComponent = SignupComponent;
//# sourceMappingURL=signup.component.js.map