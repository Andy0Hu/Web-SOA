package com.soaproject.demo.controller;

import com.soaproject.demo.api.StatusCode;
import com.soaproject.demo.api.BaseResponse;
import com.soaproject.demo.server.request.CheckRequest;
import com.soaproject.demo.server.request.OrderRequest;
import com.soaproject.demo.server.service.OrderService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.MediaType;
import org.springframework.validation.BindingResult;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import javax.validation.Valid;


@RestController
public class ProducerController {
    private static final Logger log= LoggerFactory.getLogger(ProducerController.class);

    private static final String prefix="myorders/";

    @Resource
    private OrderService orderService;

    /**
     * 创建用户下单记录
     * @param requestData
     * @param bindingResult
     * @return
     * @throws Exception
     */

    @RequestMapping(value = prefix+"/create",method = RequestMethod.POST,consumes = MediaType.APPLICATION_JSON_UTF8_VALUE,produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    public BaseResponse createRecord(@Valid @RequestBody OrderRequest requestData, BindingResult bindingResult) throws Exception{
        if (bindingResult.hasErrors()){
            bindingResult.getAllErrors().forEach(o ->{
                FieldError error = (FieldError) o;
                log.info(error.getField() + ":" + error.getDefaultMessage());
            });
            return new BaseResponse(StatusCode.PARAM_ERROR);
        }
        BaseResponse response=new BaseResponse(StatusCode.SUCCESS);
        try {
            orderService.createTradeRecord(requestData);
        }catch (Exception e){
            log.error("用户下单记录异常：{} ",requestData,e.fillInStackTrace());
            return new BaseResponse(StatusCode.FAILED);
        }
        return response;
    }

    /**
     * 显示订单记录
     * @param requestData
     * @param bindingResult
     * @return
     * @throws Exception
     */
    @GetMapping(value="/order")
    public BaseResponse checkRecord(@Valid @RequestBody CheckRequest requestData, BindingResult bindingResult) throws Exception{
        if (bindingResult.hasErrors()){
            bindingResult.getAllErrors().forEach(o ->{
                FieldError error = (FieldError) o;
                log.info(error.getField() + ":" + error.getDefaultMessage());
            });
            return new BaseResponse(StatusCode.PARAM_ERROR);
        }
        BaseResponse response=new BaseResponse(StatusCode.SUCCESS);
        try {
            orderService.checkRecord(requestData);
        }catch (Exception e){
            log.error("查询订单记录异常：{} ",requestData,e.fillInStackTrace());
            return new BaseResponse(StatusCode.FAILED);
        }
        return response;
    }


}
