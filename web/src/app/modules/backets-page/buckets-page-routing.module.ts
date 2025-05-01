import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BucketsListComponent} from './backets-list/buckets-list.component';

const routes: Routes = [
  {
    path: '',
    component: BucketsListComponent,
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BucketsPageRoutingModule { }
