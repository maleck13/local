import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions, Response } from '@angular/http';

import 'rxjs/add/operator/toPromise';

@Injectable()
export class CouncillorsService {


  constructor(private http: Http) { }

  private councillorURL = "/councillors";

  

  public councillors(county: string, area: string, auth: Headers): Promise<Councillor[]> {
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.councillorURL + "?county=" + county;
    if (area && area != "") {
      calURL += "&area=" + area
    }
    return this.http.get(calURL, options)
      .toPromise()
      .then((res) => {
        let cs = res.json().map((c) => {
            return new Councillor(c.id, c.firstName, c.secondName, c.area, c.email, c.twitter, c.facebook, c.phone, c.party, c.web, c.county, c.address,c.image)
        });
        return cs
      })
      .catch(this.handleError);
  }

  public councillor(id: string, auth: Headers): Promise<Councillor> {
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.councillorURL + "/" + id
    return this.http.get(calURL, options)
      .toPromise()
      .then((res) => {
          let c = res.json();
          return new Councillor(c.id, c.firstName, c.secondName, c.area, c.email, c.twitter, c.facebook, c.phone, c.party, c.web, c.county, c.address,c.image)
       })
      .catch(this.handleError);

  }

  public constituents(id: string, auth:Headers):Promise<Constituent[]>{
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.councillorURL + "/" + id + "/constituents"
    return this.http.get(calURL, options)
      .toPromise()
      .then((res) => {
          let cs = res.json();
          let constituents = []
          cs.each((item)=>{
            constituents.push(new Constituent(item.id,item.firstName,item.secondName));
          })
          return constituents;
       })
      .catch(this.handleError);
  }

  public update(c :Councillor, auth:Headers):Promise<Response>{
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.councillorURL + "/" + c.id;
    return this.http.post(calURL, c,options)
      .toPromise()
      .catch(this.handleError);
  }

  private handleError(error: Response) {
    return Promise.reject(error);
  }
  private handleResponse(json: any[]): Councillor[] {
    let cs = json.map((c) => {
      return new Councillor(c.id, c.firstName, c.secondName, c.area, c.email, c.twitter, c.facebook, c.phone, c.party, c.web, c.county, c.address,c.image)
    })
    return cs
  }


}


export class Councillor {
  constructor(
    public id: string,
    public firstName: string,
    public secondName: string,
    public area: string,
    public email: string,
    public twitter: string,
    public facebook: string,
    public phone: string,
    public party: string,
    public web: string,
    public county: string,
    public address: string,
    public image:string)
  { }
}

export class Constituent{
  constructor(
    public id: string,
    public firstName: string,
    public secondName: string
  ){}
}

export class CouncillorCommunication{
  constructor(
    public id:string,
    public subject:string,
    public recepientID:string,
    public body:string,
    public isPrivate:boolean,
    public type:string,
    public commID:string 
  ){}
}