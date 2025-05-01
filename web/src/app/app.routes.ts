import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: "",
    loadChildren: () =>
      import('./modules/home-page/home-page.module').then(x => x.HomePageModule)
  },
  {
    path: "buckets",
    loadChildren: () =>
      import('./modules/backets-page/buckets-page.module').then(x => x.BucketsPageModule)
  }
];
