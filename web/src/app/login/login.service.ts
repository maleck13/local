import { Injectable, EventEmitter} from '@angular/core';
import { Headers, Http, RequestOptions, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import {Observable, Observer} from 'rxjs/Rx'


@Injectable()
export class LoginService {

    public loggedIn$: Observable<Boolean>;
    private loggedInObserver: Observer<any>;
    private loginUrl = 'user/login';


    constructor(private http: Http) {
        var self = this;
        this.loggedIn$ = new Observable<Boolean>((observer) => {
            this.loggedInObserver = observer;
        }).share();
    }

    login(signup: Login): Promise<any> {
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(signup);

        return this.http.post(this.loginUrl, body, options)
            .toPromise()
            .then(this.handleLogin)
            .then((data) => {
                if (this.loggedInObserver && this.loggedInObserver.next) {
                    this.loggedInObserver.next(true);
                }
                return data;
            })
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    }
    private handleLogin(res: any): any {
        let body = res.json();
        if (res.status == 401) {
            throw new Error("bad login ");
        }
        return body || {};
    }


}

export class Login {
    constructor(
        public token: string,
        public email: string,
        public signupType: string
    ) { }
}
