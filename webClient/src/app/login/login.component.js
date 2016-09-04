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
var LoginComponent = (function () {
    function LoginComponent(_zone) {
        var _this = this;
        this._zone = _zone;
        this.userAuthToken = null;
        this.userDisplayName = "empty";
        this.userEmail = "";
        // Triggered after a user successfully logs in using the Google external
        // login provider.
        this.onGoogleLoginSuccess = function (loggedInUser) {
            console.log("here");
            _this._zone.run(function () {
                _this.userAuthToken = loggedInUser.getAuthResponse().id_token;
                _this.userDisplayName = loggedInUser.getBasicProfile().getName();
                _this.userEmail = loggedInUser.getBasicProfile().getEmail();
                console.log(_this.userDisplayName, _this.userAuthToken, _this.userEmail);
            });
        };
    }
    LoginComponent.prototype.ngAfterViewInit = function () {
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
    LoginComponent.prototype.ngOnInit = function () {
    };
    LoginComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'app-login',
            templateUrl: 'login.component.html',
            styleUrls: ['login.component.css']
        }), 
        __metadata('design:paramtypes', [core_1.NgZone])
    ], LoginComponent);
    return LoginComponent;
})();
exports.LoginComponent = LoginComponent;
//# sourceMappingURL=login.component.js.map