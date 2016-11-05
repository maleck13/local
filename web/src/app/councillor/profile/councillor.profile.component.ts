import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import {Councillor} from '../../admin/admin.councillor.service'
import {CouncillorsService} from '../../councillors/councillors.service'

@Component({
  selector: 'councillor',
  templateUrl: 'councillor.profile.component.html',
  styleUrls: ['councillor.profile.component.css'],
  providers: [CouncillorsService]
})

export class CouncillorProfileComponent implements OnInit {

        ngOnInit(){

        }

}