package com.songshu.snt.nuts.core.service;

import java.util.List;

import com.songshu.snt.nuts.client.dto.{{.Name}}DTO;
import com.songshu.snt.nuts.client.query.{{.Name}}Query;


/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
public interface {{.Name}}Service {

    /**
     * 保存数据
     * @param {{.Name | lowerCamelCase}}DTO
     * @return
     */
    Integer save({{.Name}}DTO {{.Name | lowerCamelCase}}DTO);

    /**
     * 根据ID逻辑删除数据
     * @param id
     * @return
     */
    Integer delete(Long id);

    /**
     * 修改数据,必须带有ID
     * @param {{.Name | lowerCamelCase}}DTO
     * @return
     */
    Integer update({{.Name}}DTO {{.Name | lowerCamelCase}}DTO);

    /**
     * 根据ID查询数据
     * @param id
     * @return
     */
    {{.Name}}DTO getById(Long id);

    /**
     * 根据查询条件获取数据
     * @param {{.Name | lowerCamelCase}}Query
     * @return
     */
    List<{{.Name}}DTO> query({{.Name}}Query {{.Name | lowerCamelCase}}Query);

}
