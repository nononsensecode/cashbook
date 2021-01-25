import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { AppRoutingModule } from "src/app/app-routing.module";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { SidenavComponent } from "../../components/sidenav/sidenav.component";
import { ToolbarComponent } from "../../components/toolbar/toolbar.component";

@NgModule({
    declarations: [
        SidenavComponent,
        ToolbarComponent
    ],
    imports: [
        CommonModule,
        MaterialModule,
        AppRoutingModule
    ],    
    exports: [
        SidenavComponent,
        ToolbarComponent
    ]
})
export class NavigationModule{}