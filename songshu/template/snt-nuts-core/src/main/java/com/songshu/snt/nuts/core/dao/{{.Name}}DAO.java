package com.songshu.snt.nuts.core.dao;

import java.util.List;

import org.apache.ibatis.annotations.Param;

import com.songshu.snt.base.db.SntBaseDAO;

import com.songshu.snt.nuts.core.entity.{{.Name}}DO;
import com.songshu.snt.nuts.client.query.{{.Name}}Query;


/**
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
public interface {{.Name}}DAO extends SntBaseDAO<{{.Name}}DO,{{.Name}}Query> {

    /**
     * 逻辑删除
     * @param id
     * @return
     */
    Integer delete(@Param("id")Long id);

   /**
     * 根据主键查询
     * @param id
     * @return
     */
    {{.Name}}DO queryById(@Param("id")Long id);
}
