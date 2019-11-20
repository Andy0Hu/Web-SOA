package com.soaproject.demo.controller;


import com.alibaba.fastjson.JSONObject;
import com.soaproject.demo.common.ProjectProcessConstant;
import com.soaproject.demo.dao.WaitCancelOrderMapper;
import com.soaproject.demo.entity.WaitCancelOrder;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.camunda.bpm.engine.RuntimeService;
import org.camunda.bpm.engine.TaskService;
import org.camunda.bpm.engine.runtime.ProcessInstance;
import org.camunda.bpm.engine.task.Task;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.*;

@Api(value = "取消下单")
@RestController
@RequestMapping("/cancelorder")
public class CancelOrderController {
    private final Logger LOGGER = LoggerFactory.getLogger(CancelOrderController.class);
    private final TaskService taskService;
    private final RuntimeService runtimeService;

    @Resource
    public WaitCancelOrderMapper waitCancelOrderMapper;

    public CancelOrderController(TaskService taskService, RuntimeService runtimeService) {
        this.taskService = taskService;
        this.runtimeService = runtimeService;
    }

    //发起申请
    @ApiOperation(value = "申请取消订单")
    @PostMapping(value = "/{corderId}/users/{cuserId}")
    public boolean ParticipatingProject(@PathVariable String corderId, @PathVariable String cuserId) {
        //ignore argument verify

        //save the record to db
        Random r = new Random();
        Long savedCTaskId = r.nextLong();
        WaitCancelOrder waitCancelOrder=new WaitCancelOrder();
        waitCancelOrder.setCtaskId(String.valueOf(savedCTaskId));
        waitCancelOrder.setCorderId(corderId);
        waitCancelOrder.setCuserId(cuserId);
        waitCancelOrderMapper.insertSelective(waitCancelOrder);

        //start a new instance of the process
        Map<String, Object> variables = new HashMap<String, Object>();
        variables.put(ProjectProcessConstant.VAR_CUSER, cuserId);
        variables.put(ProjectProcessConstant.VAR_CORDER, corderId);
        variables.put(ProjectProcessConstant.VAR_CTASK, String.valueOf(savedCTaskId));

        ProcessInstance instance = runtimeService.
                startProcessInstanceByKey(ProjectProcessConstant.PROCESS_ID, variables);
        if (instance == null) {
            return false;
        }else {
            return true;
        }
    }

    @ApiOperation(value = "获取需要审批的项目申请列表")
    @GetMapping(value = "/project/approve/list")
    public @ResponseBody List<WaitCancelOrder> getAllWaitCancelOrder() {

        //get the taskList
        List<Task> tasks;
        tasks = taskService.createTaskQuery().
                taskName(ProjectProcessConstant.TASK_NAME_REVIEW).
                list();

        List<WaitCancelOrder> records = new ArrayList<WaitCancelOrder>(tasks.size());
        tasks.forEach( task -> {
            WaitCancelOrder record = new WaitCancelOrder();
            String taskId = task.getId();
            Map<String, Object> variables = taskService.getVariables(taskId);

            String cTaskId = (String)variables.get(ProjectProcessConstant.VAR_CTASK) ;
            String cUserId = (String)variables.get(ProjectProcessConstant.VAR_CUSER);
            String cOrderId = (String)variables.get(ProjectProcessConstant.VAR_CORDER);
            record.setCtaskId(cTaskId);
            record.setCorderId(cOrderId);
            record.setCuserId(cUserId);
            record.setTaskId(taskId);
            waitCancelOrderMapper.updateByPrimaryKeySelective(record);
            records.add(record);
        });

        return records;
    }

    //审批
    @ApiOperation(value = "审批取消申请")
    @PutMapping(value = "/project/participateRequests/{taskId}")
    public boolean approveProjectParticipateRequest(@PathVariable String taskId, boolean passed) {
        Task task = taskService.createTaskQuery().
                taskId(taskId).singleResult();
        if (task == null) {
            return false;
        }else {
            //business logic here

            //Into next step
            Map<String, Object> variables = new HashMap<>();
            variables.put(ProjectProcessConstant.FORM_APPROVED_1, passed);
            taskService.complete(task.getId(), variables);
            return true;
        }
    }





}
