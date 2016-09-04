import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';


@Injectable()
export class LoginService {

  constructor(private http: Http) { }
    private loginUrl = 'user/login';

    login(signup:Login):Promise{
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(signup);

        return this.http.post(this.loginUrl,body,options)
            .toPromise()
            .then(this.extractData)
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    }
    private extractData(res: Response) {
        let body = res.json();
        
        return body || { };
    }   

}



export class Login{
    constructor(
        public token: string,
        public email: string,
        public signupType: string
    ){}
}
