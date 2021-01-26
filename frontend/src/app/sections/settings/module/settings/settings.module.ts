import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { SettingsMainComponent } from "../../components/settings-main/settings-main.component";

@NgModule({
    declarations:[
        SettingsMainComponent
    ],
    imports: [
        CommonModule,
        MaterialModule
    ],
    exports: [
        SettingsMainComponent
    ]
})
export class SettingsModule{}