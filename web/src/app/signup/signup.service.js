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
var http_1 = require('@angular/http');
require('rxjs/add/operator/toPromise');
var SignupService = (function () {
    function SignupService(http) {
        this.http = http;
        this.signUpUrl = 'api/signup';
    }
    SignupService.prototype.signUp = function (signup) {
        var headers = new http_1.Headers({ 'Content-Type': 'application/json' });
        var options = new http_1.RequestOptions({ headers: headers });
        var body = JSON.stringify(signup);
        return this.http.post(this.signUpUrl, body, options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    SignupService.prototype.handleError = function (error) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    };
    SignupService = __decorate([
        core_1.Injectable(), 
        __metadata('design:paramtypes', [http_1.Http])
    ], SignupService);
    return SignupService;
})();
exports.SignupService = SignupService;
var Signup = (function () {
    function Signup(signUpType, token, email, firstName, secondName) {
        this.signUpType = signUpType;
        this.token = token;
        this.email = email;
        this.firstName = firstName;
        this.secondName = secondName;
    }
    return Signup;
})();
exports.Signup = Signup;
//# sourceMappingURL=signup.service.js.map