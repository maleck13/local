import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions } from '@angular/http';

import 'rxjs/add/operator/toPromise';

@Injectable()
export class SignupService {
    constructor(private http: Http) { }
    private signUpUrl = 'user/signup';

    signUp(signup:Signup):Promise{
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(signup);

        return this.http.post(this.signUpUrl,body,options)
            .toPromise()
            .then((res)=>res)
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    }

}

export class Signup{
    constructor(
        public signUpType: string,
        public token: string,
        public email: string,
        public firstName: string,
        public secondName: string
    ){}
}
