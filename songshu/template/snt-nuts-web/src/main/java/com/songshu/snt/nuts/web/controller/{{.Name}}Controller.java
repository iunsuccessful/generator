package com.songshu.snt.nuts.web.controller;

import com.songshu.snt.base.common.SntResult;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import com.songshu.snt.base.common.query.SntPage;
import com.songshu.snt.base.common.err.ErrorCodeSourceEnum;
import com.songshu.snt.base.common.err.ErrorCodeTypeEnum;
import com.songshu.snt.base.common.exception.SntException;
import com.songshu.snt.base.common.exception.UnknownException;
import java.util.List;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import com.songshu.snt.nuts.client.dto.{{.Name}}DTO;
import com.songshu.snt.nuts.client.query.{{.Name}}Query;
import com.songshu.snt.nuts.core.service.{{.Name}}Service;


/**
 *
 * TODO 入参校验
 *
 * <pre>
 * Author 依韵 {{datetime}}
 * </pre>
 */
@RestController
public class {{.Name}}Controller {

    private static Logger LOGGER = LoggerFactory.getLogger({{.Name}}Controller.class);

    @Autowired
    private {{.Name}}Service {{.Name | lowerCamelCase}}Service;


    /**
     * post一个对象入库
     * @param {{.Name | lowerCamelCase}}DTO
     * @return
     */
    @PostMapping("/{{.Name | lowerCamelCase}}")
    public SntResult<Integer> save{{.Name}}(@RequestBody {{.Name}}DTO {{.Name | lowerCamelCase}}DTO){
        try{
            if(null == {{.Name | lowerCamelCase}}DTO){
                return SntResult.fail("errorCode","入参错误");
            }

            Integer cnt = {{.Name | lowerCamelCase}}Service.save({{.Name | lowerCamelCase}}DTO);
            //如果cnt为0不符合业务，需要自己判断后设置result为fail
            return SntResult.ok(cnt);
        }catch (SntException e){
            LOGGER.error("save{{.Name}} SntException,{}",{{.Name | lowerCamelCase}}DTO,e);
            return SntResult.fail(e);
        }catch (Exception e){
            LOGGER.error("save{{.Name}} exception,{}",{{.Name | lowerCamelCase}}DTO,e);
            return SntResult.fail(new UnknownException(e));
        }
    }


    /**
    * 删除一个对象
    * @return
    */
    @DeleteMapping("/{{.Name | lowerCamelCase}}/{id}")
    public SntResult<Integer> delete{{.Name}}(@PathVariable Long id){
        try{
            if(null == id || id < 1){
                return SntResult.fail("errorCode","入参错误");
            }

            Integer cnt = {{.Name | lowerCamelCase}}Service.delete(id);
            //如果cnt为0不符合业务，需要自己判断后设置result为fail
            return SntResult.ok(cnt);
        }catch (SntException e){
            LOGGER.error("delete{{.Name}} SntException,{}",id,e);
            return SntResult.fail(e);
        }catch (Exception e){
            LOGGER.error("delete{{.Name}} exception,{}",id,e);
            return SntResult.fail(new UnknownException(e));
        }
    }

    /**
    * 修改一个对象
    * @param {{.Name | lowerCamelCase}}DTO
    * @return
    */
    @PutMapping("/{{.Name | lowerCamelCase}}")
    public SntResult<Integer> update{{.Name}}(@RequestBody {{.Name}}DTO {{.Name | lowerCamelCase}}DTO){
        try{
            if(null == {{.Name | lowerCamelCase}}DTO || {{.Name | lowerCamelCase}}DTO.getId() == null){
                return SntResult.fail("errorCode","入参错误");
            }

            Integer cnt = {{.Name | lowerCamelCase}}Service.update({{.Name | lowerCamelCase}}DTO);
            //如果cnt为0不符合业务，需要自己判断后设置result为fail
            return SntResult.ok(cnt);
        }catch (SntException e){
            LOGGER.error("update{{.Name}} SntException,{}",{{.Name | lowerCamelCase}}DTO,e);
            return SntResult.fail(e);
        }catch (Exception e){
            LOGGER.error("update{{.Name}} exception,{}",{{.Name | lowerCamelCase}}DTO,e);
            return SntResult.fail(new UnknownException(e));
        }
    }


    /**
    * 根据ID查询一个对象
    * @param id
    * @return
    */
    @GetMapping("/{{.Name | lowerCamelCase}}/{id}")
    public SntResult<{{.Name}}DTO> get{{.Name}}ById(@PathVariable Long id){
        try{
            if(null == id || id < 1){
                return SntResult.fail("errorCode","入参错误");
            }

            {{.Name}}DTO {{.Name | lowerCamelCase}}DTO = {{.Name | lowerCamelCase}}Service.getById(id);
            return SntResult.ok({{.Name | lowerCamelCase}}DTO);
        }catch (SntException e){
            LOGGER.error("get{{.Name}}ById SntException,{}",id,e);
            return SntResult.fail(e);
        }catch (Exception e){
            LOGGER.error("get{{.Name}}ById exception,{}",id,e);
            return SntResult.fail(new UnknownException(e));
        }
    }

    /**
    *
    * @param {{.Name | lowerCamelCase}}Query 查询条件，根据业务自行扩展
    * @return
    */
    @GetMapping("/{{.Name | lowerCamelCase}}/query")
    public SntResult<List<{{.Name}}DTO>> query{{.Name}}({{.Name}}Query {{.Name | lowerCamelCase}}Query ){
        try{
            if(null == {{.Name | lowerCamelCase}}Query){
                return SntResult.fail("errorCode","入参错误");
            }
            //分页
            List<{{.Name}}DTO> list = {{.Name | lowerCamelCase}}Service.query({{.Name | lowerCamelCase}}Query);
            return SntResult.ok(list,{{.Name | lowerCamelCase}}Query);
        }catch (SntException e){
            LOGGER.error("query{{.Name}} SntException,{}",{{.Name | lowerCamelCase}}Query,e);
            return SntResult.fail(e);
        }catch (Exception e){
            LOGGER.error("query{{.Name}} exception,{}",{{.Name | lowerCamelCase}}Query,e);
            return SntResult.fail(new UnknownException(e));
        }
    }
}
