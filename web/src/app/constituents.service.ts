import { Injectable } from '@angular/core';
import {Http, Headers,Response,RequestOptions} from '@angular/http'
import 'rxjs/add/operator/toPromise';

@Injectable()
export class ConstituentsService {

  constructor(private http:Http) { }

  public constituents(cid:String , auth:Headers):Promise<Constituent[]>{
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let url = "councillors/"+cid+"/consituents";
    return this.http.get(url,options).toPromise()
    .then((res)=>{
      let cs = res.json()
      if (Array.isArray(cs)){
        let consituents = cs.map((c)=>{
          return new Constituent(c.id,c.firstName,c.secondName,c.openComms,c.hasOpenComms);
        })
        return consituents;
      }
    })
  }

}

export class Constituent{
  constructor(public id:String,public firstName:String,public secondName:String, public openComms:any[],public hasOpenComms:Boolean){}
}
