import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { DashboardMainComponent } from './../../components/dashboard-main/dashboard-main.component';

@NgModule({
    declarations:[
        DashboardMainComponent
    ],
    imports: [
        CommonModule,
        MaterialModule
    ],
    exports: [
        DashboardMainComponent
    ]
})
export class DashboardModule {}