package com.songshu.snt.nuts.client.query;

import com.songshu.snt.base.common.query.BaseQuery;
import lombok.Data;
import java.util.Date;


/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
@Data
public class {{.Name}}Query extends BaseQuery {

    private static final long serialVersionUID = 1L;
    
    {{range .Columns}}
    /**
     * {{.ColumnComment}}
     */
    private {{.DataType}} {{.ColumnName | lowerCamelCase}};
    {{end}}


    public {{.Name}}Query addParamSort(String column, String sort){
        this.setColumn(column);
        this.setSort(sort);
        return this;
    }

    public {{.Name}}Query addParamPageSize(Integer pageSize){
        this.setPageSize(pageSize);
        return this;
    }
}
