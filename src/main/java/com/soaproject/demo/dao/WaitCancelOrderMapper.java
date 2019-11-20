package com.soaproject.demo.dao;

import com.soaproject.demo.entity.WaitCancelOrder;
import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Component;

@Component
@Mapper
public interface WaitCancelOrderMapper {
    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    int deleteByPrimaryKey(String ctaskId);

    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    int insert(WaitCancelOrder record);

    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    int insertSelective(WaitCancelOrder record);

    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    WaitCancelOrder selectByPrimaryKey(String ctaskId);

    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    int updateByPrimaryKeySelective(WaitCancelOrder record);

    /**
     * This method was generated by MyBatis Generator.
     * This method corresponds to the database table waitcancelorder
     *
     * @mbggenerated
     */
    int updateByPrimaryKey(WaitCancelOrder record);


    WaitCancelOrder selectAll();
}