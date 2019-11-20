package com.soaproject.demo.server.service;

import com.soaproject.demo.dao.OrderMapper;
import com.soaproject.demo.dao.WaitCancelOrderMapper;
import com.soaproject.demo.entity.Order;
import org.camunda.bpm.engine.TaskService;
import org.camunda.bpm.engine.delegate.DelegateExecution;
import org.camunda.bpm.engine.delegate.JavaDelegate;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Map;

@Component
@Service
public class SmsServiceTask implements JavaDelegate {
    private static final Logger log = LoggerFactory.getLogger(SmsServiceTask.class);

    private final TaskService taskService;

    @Resource
    public OrderMapper orderMapper;

    @Resource
    public WaitCancelOrderMapper waitCancelOrderMapper;

    public SmsServiceTask(TaskService taskService) {
        this.taskService = taskService;
    }

    public void execute(DelegateExecution delegateExecution) throws Exception {
        Map<String, Object> variables = delegateExecution.getVariables();
        log.info("variables is {}", variables);

        String corderID = (String)variables.get("corderID");
        String ctaskID = (String)variables.get("ctaskID");
        waitCancelOrderMapper.deleteByPrimaryKey(ctaskID);
        orderMapper.deleteByPrimaryKey(corderID);
    }
}
