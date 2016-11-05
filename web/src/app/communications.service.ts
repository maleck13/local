import { Injectable } from '@angular/core';
import {Http, Headers,Response,RequestOptions} from '@angular/http'
import 'rxjs/add/operator/toPromise';
import {CouncillorCommunication} from './councillors/councillors.service'
@Injectable()
export class CommunicationsService {

  constructor(private http:Http) { }

  private communicationURL = "/communications"

  communicate(com:CouncillorCommunication, auth:Headers):Promise<Response>{
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.communicationURL + "/send"
    return this.http.post(calURL,com,options).toPromise()
    .then((res)=>{
      return res.json()
    })
    .catch((err)=>this.handleError);
  }

  listForUser(cid:string, auth:Headers,commID:string):Promise<CouncillorCommunication[]>{
    auth.append('Content-Type', 'application/json');
    let options = new RequestOptions({ headers: auth });
    let calURL = this.communicationURL + "/councillor/" + cid
    if (commID){
      calURL +="?commID="+ commID
    }
    return this.http.get(calURL,options).toPromise()
    .then((res)=>{
      let  comms =  res.json()
      if (Array.isArray(comms)){
        let communications = comms.map((c)=>{
          return new CouncillorCommunication(c.id,c.subject,c.recepientID,c.body,c.isPrivate,"email",c.commID)
        })
        return communications;
      }
    })
    .catch((err)=>this.handleError);
  }

  private handleError(error: Response) {
    return Promise.reject(error);
  }

}
