package com.soaproject.demo.server.service;

import com.alibaba.fastjson.JSONObject;
import com.soaproject.demo.dao.OrderMapper;
import com.soaproject.demo.entity.Order;
import com.soaproject.demo.server.request.CheckRequest;
import com.soaproject.demo.server.request.OrderRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.support.converter.Jackson2JsonMessageConverter;
import org.springframework.beans.BeanUtils;
import org.springframework.core.env.Environment;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
public class OrderService {
    private static final Logger log= LoggerFactory.getLogger(OrderService.class);

    @Resource
    private Environment env;

    @Resource
    public OrderMapper orderMapper;

    @Resource
    private RabbitTemplate rabbitTemplate;

    @Resource
    private StringRedisTemplate redisTemplate;
    public static final String ORDER_KEY="ORDER";

    public void createTradeRecord(OrderRequest requestData) throws Exception{
        //TODO:其余业务逻辑上的校验。。

        //TODO：创建交易记录
        Order order=new Order();
        BeanUtils.copyProperties(requestData,order);
        order.setOrderStatus("未完成");
        order.setDeliveryInfo("发货中");
        String jsonString = JSONObject.toJSONString(order);
        redisTemplate.opsForValue().set("ORDER",jsonString);
        orderMapper.insertSelective(order);

        //TODO：设置超时，用mq处理已超时的下单记录（一旦记录超时，则处理为无效）
        final Long ttl=env.getProperty("trade.record.ttl",Long.class);
        rabbitTemplate.setMessageConverter(new Jackson2JsonMessageConverter());
        rabbitTemplate.setExchange(env.getProperty("register.delay.exchange.name"));
        rabbitTemplate.setRoutingKey("");
        rabbitTemplate.convertAndSend(order, message -> {
            message.getMessageProperties().setExpiration(ttl+"");
            return message;
        });
    }

    public Order checkRecord(CheckRequest requestData) throws Exception{

        Order retOrder = orderMapper.selectByPrimaryKey(requestData.getOrderId());
        return retOrder;
    }
}
