import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';


export class Councillor{
    constructor(
        public firstName:string,
        public secondName:string,
        public area:string,
        public email:string,
        public twitter:string,
        public facebook:string,
        public phone:string,
        public party:string,
        public web:string,
        public county:string,
        public address:string)
        {}
}