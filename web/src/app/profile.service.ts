import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions } from '@angular/http';
import {Observer,Observable} from 'rxjs/RX'

export interface AuthenticatedService{
  getTokenHeader: any
}

@Injectable()
export class ProfileService implements AuthenticatedService{

  constructor(private http:Http) {
    this.userDataObservable = new Observable<UserData>((s)=> this.userDataObserver = s)
  }

  profileUrl: string = "/user";

  userDataObserver:Observer<UserData>
  userDataObservable:Observable<UserData>

  

  public storeUserData(userData : UserData):Promise<any>{
      let prom = new Promise((res,rej)=>{
        if(localStorage){
        localStorage.setItem("profile", JSON.stringify(userData))
        if (this.userDataObserver && this.userDataObserver.next){
          this.userDataObserver.next(userData);
        }
        res();
      }else{
        //todo store in a cookie
        rej(new Error("failed to store user data no localStorage"));
      }
      });
      return prom;
  }

  public getUserData():Promise<UserData>{
    let prom = new Promise((res,rej)=>{
      if(localStorage){
        let data = localStorage.getItem("profile");
        if(! data)return res();
        let user:UserData ;
        try{
          let u = JSON.parse(data); 
          user = new UserData(u.id,u.token,u.type,u.county,u.area);
          return res(user);
        }catch(e){
          return rej(e);
        }
      }else{
        rej(new Error("no local storage available"));
      }
    });
    return prom;
  }

  public deleteUserData():Promise<any>{
    let prom = new Promise((res,rej)=>{
      if(localStorage){
        localStorage.clear();
        res();
      }else{
        rej(new Error("no local storage available"));
      }
    });
    return prom;
  }

  public getTokenHeader():Promise<Headers>{
    let p = new Promise((res,rej)=>{
      this.getUserData().then((d)=>{
        if(d){
          return res(new Headers({"Authorization":"  Bearer " + d.token}));
        }
        return res();
      }).catch((err)=>{
          return rej(err);
      })
    });
    return p;
  }

  public getJwtToken():Promise<String>{
    let p = new Promise((res,rej)=>{
      this.getUserData().then((d)=>{
        if(d){
          return res(d.token);
        }
        return res();
      }).catch((err)=>{
         return rej(err);
      })
    });
      return p;
  }

  public getProfile(id : string):Promise<any>{
    return this.getTokenHeader().then((header)=>{
      let headers = new Headers(header);
      headers.append("Content-type","application/json");
      let options = new RequestOptions({ headers: headers });
        return this.http.get(this.profileUrl + "/" + id ,options)
            .toPromise()
            .then((res)=>res.json())
            .catch(this.handleError);
    }); 
  }

  public updateProfile(profile: Profile):Promise<any>{
    return this.getTokenHeader().then((header)=>{
      let headers = new Headers(header);
      headers.append("Content-type","application/json");
      let options = new RequestOptions({ headers: headers });
        return this.http.post(this.profileUrl + "/" + profile.id, profile ,options)
            .toPromise()
            .then((res)=>res.json())
            .catch(this.handleError);
    }); 
  }

  private handleError(error: any) {
        console.error('An error occurred in signup', error);
        return Promise.reject(error);
    }

}


export class UserData{
  constructor(public id: string , public token:string, public type:string, public county:string, public area:string){}
}

export class Profile {
  constructor(
  public id:string,  
  public county:string,
  public area:string,
  public email:string,
  public firstName:string,
  public secondName:string){}
}
