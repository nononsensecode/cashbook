import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { FooterComponent } from "../../components/footer/footer.component";
import { HeaderComponent } from "../../components/header/header.component";

@NgModule({
    declarations: [
        HeaderComponent,
        FooterComponent
    ],
    imports: [
        CommonModule,
        MaterialModule
    ],
    exports: [
        HeaderComponent,
        FooterComponent
    ]
})
export class HeaderFooterModule {}