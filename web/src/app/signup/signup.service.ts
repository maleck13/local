import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions , Response} from '@angular/http';

import 'rxjs/add/operator/toPromise';

@Injectable()
export class SignupService {
    constructor(private http: Http) { }
    private signUpUrl = 'user/signup';
    private checkCouncillorExistsURL = "user/councillor/signup" 

    signUp(signup:Signup):Promise<any>{
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(signup);

        return this.http.post(this.signUpUrl,body,options)
            .toPromise()
            .then((res)=>res)
            .catch(this.handleError);
    }

    councillorUserExists(email:string):Promise<Response>{
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify({"email":email});

        return this.http.post(this.checkCouncillorExistsURL,body,options)
            .toPromise()
            .then((res)=>{
                return res 
            })
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
        public secondName: string,
        public type: string,
        public county: string
    ){}
}
