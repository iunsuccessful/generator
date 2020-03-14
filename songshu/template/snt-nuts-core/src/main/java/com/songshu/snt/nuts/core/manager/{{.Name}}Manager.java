package com.songshu.snt.nuts.core.manager;

import org.springframework.stereotype.Component;
import org.springframework.util.Assert;
import javax.annotation.Resource;
import java.util.List;

import com.songshu.snt.nuts.core.entity.{{.Name}}DO;
import com.songshu.snt.nuts.client.query.{{.Name}}Query;
import com.songshu.snt.nuts.core.dao.{{.Name}}DAO;

/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
@Component
public class {{.Name}}Manager {

    @Resource
    private {{.Name}}DAO {{.Name | lowerCamelCase}}DAO;

    public Integer save({{.Name}}DO {{.Name | lowerCamelCase}}DO) {
        Integer cnt = {{.Name | lowerCamelCase}}DAO.insert({{.Name | lowerCamelCase}}DO);
        return cnt;
    }

    public Integer delete(Long id) {
        Assert.notNull(id, "id cant null");
        Integer cnt = {{.Name | lowerCamelCase}}DAO.delete(id);
        return cnt;
    }

    public Integer update({{.Name}}DO {{.Name | lowerCamelCase}}DO) {
        Integer cnt = {{.Name | lowerCamelCase}}DAO.update({{.Name | lowerCamelCase}}DO);
        return cnt;
    }

    public {{.Name}}DO getById(Long id) {
        {{.Name}}DO {{.Name | lowerCamelCase}}DO = {{.Name | lowerCamelCase}}DAO.queryById(id);
        return {{.Name | lowerCamelCase}}DO;
    }

    public List<{{.Name}}DO> query({{.Name}}Query {{.Name | lowerCamelCase}}Query) {
        List<{{.Name}}DO> {{.Name | lowerCamelCase}}List = {{.Name | lowerCamelCase}}DAO.query({{.Name | lowerCamelCase}}Query);
        return {{.Name | lowerCamelCase}}List;
    }

    public Integer queryCount({{.Name}}Query {{.Name | lowerCamelCase}}Query) {
        Integer cnt = {{.Name | lowerCamelCase}}DAO.queryCount({{.Name | lowerCamelCase}}Query);
        return cnt;
    }
}
