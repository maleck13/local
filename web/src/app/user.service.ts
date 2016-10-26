import { Injectable } from '@angular/core';
import { Headers, Http,Response, RequestOptions } from '@angular/http';
import {Observer,Observable} from 'rxjs/RX'

@Injectable()
export class UserService {

  constructor(private http:Http) { }

  private userURL = "/user"

  public resetPassword(authToken:String, newPass: String):Promise<Response>{
    let authHeader =  new Headers({"Authorization":"  Bearer " + authToken});
    authHeader.append("Content-type","application/json");
    let options = new RequestOptions({ headers: authHeader });
        return this.http.post(this.userURL + "/resetpassword",{"newpassword":newPass},options)
            .toPromise()
            .then((res)=>res)
            .catch(this.handleError);

  }

  private handleError(error: any) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    }

}
