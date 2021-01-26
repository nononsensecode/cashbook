import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AccountCodesMainComponent } from './sections/account-codes/components/account-codes-main.component';
import { CashBookEntryMainCompoent } from './sections/cash-book-entry/components/cash-book-entry-main/cash-book-entry-main.component';
import { DashboardMainComponent } from './sections/dashboard/components/dashboard-main/dashboard-main.component';
import { SettingsMainComponent } from './sections/settings/components/settings-main/settings-main.component';

const routes: Routes = [
  {
    path: 'dashboard',
    component: DashboardMainComponent
  },
  {
    path: 'cash-book-entry',
    component: CashBookEntryMainCompoent
  },
  {
    path: 'settings',
    component: SettingsMainComponent
  },
  {
    path: 'account-codes',
    component: AccountCodesMainComponent
  }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes,{useHash:true})
  ],
  exports: [RouterModule]
})

export class AppRoutingModule { }
