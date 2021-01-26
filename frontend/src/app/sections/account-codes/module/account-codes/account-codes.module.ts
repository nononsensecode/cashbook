import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { MaterialModule } from "src/app/shared/modules/material/material.module";
import { AccountCodesMainComponent } from "../../components/account-codes-main.component";

@NgModule({
    declarations: [
        AccountCodesMainComponent
    ],
    imports: [
        CommonModule,
        MaterialModule
    ],
    exports: [
        AccountCodesMainComponent
    ]
})
export class AccountCodesModule {}