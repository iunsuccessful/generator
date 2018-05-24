package com.jxs.entity;

import java.util.Date;

public class {{.Name}} {

    {{range .Columns}}
    /**
     * {{.ColumnComment}}
     */
    private {{.DataType}} {{.ColumnName | lowerCamelCase}};
    {{end}}

    {{range .Columns}}
    public {{.DataType}} get{{.ColumnName}}() {
        return {{.ColumnName | lowerCamelCase}};
    }

    public void set{{.ColumnName}}({{.DataType}} {{.ColumnName | lowerCamelCase}}) {
        this.{{.ColumnName | lowerCamelCase}} = {{.ColumnName | lowerCamelCase}};
    }
    {{end}}
    

}
