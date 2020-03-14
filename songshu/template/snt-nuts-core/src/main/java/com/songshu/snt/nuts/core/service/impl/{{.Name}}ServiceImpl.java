package com.songshu.snt.nuts.core.service.impl;

import org.springframework.stereotype.Service;
import org.springframework.beans.BeanUtils;
import javax.annotation.Resource;
import java.util.Collections;
import java.util.List;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.stream.Collectors;
import org.springframework.util.CollectionUtils;

import com.songshu.snt.nuts.core.entity.{{.Name}}DO;
import com.songshu.snt.nuts.client.dto.{{.Name}}DTO;
import com.songshu.snt.nuts.client.query.{{.Name}}Query;
import com.songshu.snt.nuts.core.manager.{{.Name}}Manager;
import com.songshu.snt.nuts.core.service.{{.Name}}Service;


/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
@Service
public class {{.Name}}ServiceImpl implements {{.Name}}Service {

    @Resource
    private {{.Name}}Manager {{.Name | lowerCamelCase}}Manager;

    public Integer save({{.Name}}DTO {{.Name | lowerCamelCase}}DTO) {
        //对象转换
        {{.Name}}DO {{.Name | lowerCamelCase}}DO = new {{.Name}}DO();
        BeanUtils.copyProperties({{.Name | lowerCamelCase}}DTO,{{.Name | lowerCamelCase}}DO);

        return {{.Name | lowerCamelCase}}Manager.save({{.Name | lowerCamelCase}}DO);

    }

    public Integer update({{.Name}}DTO {{.Name | lowerCamelCase}}DTO) {
        {{.Name}}DO {{.Name | lowerCamelCase}}DO = {{.Name | lowerCamelCase}}Manager.getById({{.Name | lowerCamelCase}}DTO.getId());
        if({{.Name | lowerCamelCase}}DO == null) {
            return 0;
        }

        //复制要修改对象的属性
        BeanUtils.copyProperties({{.Name | lowerCamelCase}}DTO, {{.Name | lowerCamelCase}}DO);

        return {{.Name | lowerCamelCase}}Manager.update({{.Name | lowerCamelCase}}DO);
    }

    public Integer delete(Long id) {
        return {{.Name | lowerCamelCase}}Manager.delete(id);
    }

    public {{.Name}}DTO getById(Long id) {

        {{.Name}}DO {{.Name | lowerCamelCase}}DO = {{.Name | lowerCamelCase}}Manager.getById(id);

        if(null == {{.Name | lowerCamelCase}}DO){
            return null;
        }

        //对象转换
        {{.Name}}DTO {{.Name | lowerCamelCase}}DTO = new {{.Name}}DTO();
        BeanUtils.copyProperties({{.Name | lowerCamelCase}}DO,{{.Name | lowerCamelCase}}DTO);

        return {{.Name | lowerCamelCase}}DTO;
    }

    public List<{{.Name}}DTO> query({{.Name}}Query {{.Name | lowerCamelCase}}Query) {
        List<{{.Name}}DO> list = {{.Name | lowerCamelCase}}Manager.query({{.Name | lowerCamelCase}}Query);

        if(CollectionUtils.isEmpty(list)){
            return Collections.emptyList();
        }

        if(null == {{.Name | lowerCamelCase}}Query.getTotalRecored()){
            Integer cnt = {{.Name | lowerCamelCase}}Manager.queryCount({{.Name | lowerCamelCase}}Query);
            {{.Name | lowerCamelCase}}Query.setTotalRecored(cnt == null ? 0 : cnt);
        }

        return list.stream()
                .map(d -> {
                    {{.Name}}DTO dto = new {{.Name}}DTO();
                    BeanUtils.copyProperties(d, dto);
                    return dto;
                })
                .collect(Collectors.toList());
    }

}
