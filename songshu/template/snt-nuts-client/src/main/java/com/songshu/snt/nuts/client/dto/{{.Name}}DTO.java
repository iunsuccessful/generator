package com.songshu.snt.nuts.client.dto;

import com.songshu.snt.base.common.DTO.SntBaseDTO;
import lombok.Data;
import java.util.Date;


/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
@Data
public class {{.Name}}DTO extends SntBaseDTO {

    private static final long serialVersionUID = 1L;
    
	{{range .Columns}}
    {{- if or (eq .ColumnName "id") (eq .ColumnName "gmt_create") (eq .ColumnName "gmt_modified") (eq .ColumnName "row_version") (eq .ColumnName "row_status") (eq .ColumnName "biz_type") (eq .ColumnName "ext_att") | not }}
    /**
     * {{.ColumnComment}}
     */
    private {{.DataType}} {{.ColumnName | lowerCamelCase}};
    {{end -}}
    {{end}}

}
