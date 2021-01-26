import { SelectionModel } from "@angular/cdk/collections";
import { Component } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { MatTableDataSource } from "@angular/material";
import { AccountCode } from "src/app/shared/domain/model/account-code";
import { HeadingService } from "src/app/shared/services/heading/heading.service";

@Component({
    selector: 'app-account-codes-main',
    templateUrl: './account-codes-main.component.html',
    styleUrls: ['./account-codes-main.component.css']
})
export class AccountCodesMainComponent {

    dataSource = new MatTableDataSource(ELEMENT_DATA);
    displayedColumns = ['select', 'id', 'code', 'description']
    selection = new SelectionModel<AccountCode>(false, []);
    public readonly addCodeForm = new FormGroup({
        id: new FormControl(''),
        code: new FormControl('', [Validators.required]),
        description: new FormControl('')
    });

    constructor(private headingService: HeadingService) {
        this.headingService.setHeading('Account Codes');
    }

    saveCode() {
        const id = (this.addCodeForm.get('id').value) ? this.addCodeForm.get('id').value : null;
        const accountCode = this.addCodeForm.get('code').value as string;
        const description = this.addCodeForm.get('description').value;
        const codeObj = {id, accountCode, description};
        if (id == null) {
            // @ts-ignore
            window.backend.createAccountCode(codeObj).then(created => {
                this.addNewCode(created);
                this.resetForm();
                this.selection.clear();
            });
        } else {
            // @ts-ignore
            window.backend.updateAccountCode(codeObj).then(updated => {
                this.updateCode(updated);
                this.resetForm();
                this.selection.clear();
            });
        }
        
    }

    reset() {
        this.resetForm();
        this.selection.clear();
    }

    private doesCodeExist(codeObj: AccountCode): boolean {
        let result = this.dataSource.data.find(code => code.id === codeObj.id);
        console.log(result);
        return result !== undefined;
    }

    private updateCode(codeObj: AccountCode) {
        const code = this.dataSource.data.find(code => code.id === codeObj.id);
        code.accountCode = codeObj.accountCode;
        code.description = codeObj.description;
    }

    private addNewCode(codeObj: AccountCode) {
        this.dataSource.data.push(codeObj);
    }

    select(event, row: AccountCode) {
        if (event.checked) {
            this.addCodeForm.controls['id'].setValue(row.id);
            this.addCodeForm.controls['code'].setValue(row.accountCode);
            this.addCodeForm.controls['description'].setValue(row.description);
        } else {
            this.resetForm();
        }
    }

    resetForm() {
        this.addCodeForm.reset();
        this.addCodeForm.get('code').setErrors(null);
    }
}

const ELEMENT_DATA: AccountCode[] = [
    {id: 1, accountCode: "1234", description: "Test code"},
    {id: 2, accountCode: "4578", description: "Test code"},
    {id: 3, accountCode: "6598", description: "Test code"},
    {id: 4, accountCode: "5879", description: "Test code"},
    {id: 5, accountCode: "7890", description: "Test code"},
];