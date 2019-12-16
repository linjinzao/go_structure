<?php
class Pd_CustomerReportRightsControlController extends Ec_Controller_Action
{
    public function preDispatch()
    {
        $this->tplDirectory = "pd/views/";
        $this->serviceClass = new Service_PdCustomerReportRightsControl();
    }
    public function getRightsByJsonAction()
    {
        $result = array('state' => 0, 'message' => 'Fail', 'data' => array());
        $paramId = $this->_request->getParam('report_id', '');
        if (!empty($paramId) && $rows = $this->serviceClass->getByCondition(array('report_id'=>$paramId),array('user_id'))) {
            $result = array('state' => 1, 'message' => '', 'data' => $rows);
        }
        die(Zend_Json::encode($result));
    }
    public function setRightsAction()
    {
        if ($this->_request->isPost()) {
            $return = array(
                "state" => 0,
                "message" => array(Ec::lang('操作失败'))
            );
            $report_id = $this ->_request->getParam('report_id','');
            $user_ids = $this ->_request->getParam('user_id',array());
            if( ''==$report_id ||  !Service_PdCustomerReport::getByField($report_id,'report_id',array('report_id'))){
                $return['message'] = array(Ec::lang('参数有误'));
            }else{
                $db_apa = Common_Common::getAdapter();
                try{
                    $db_apa->beginTransaction();
                    Service_PdCustomerReportRightsControl::delete($report_id,'report_id');
                    if(!empty($user_ids)){
                        foreach ($user_ids as $key_ui  => $val_ui){
                            if(!empty($val_ui)){
                                Service_PdCustomerReportRightsControl::add(array(
                                    'report_id' => $report_id,
                                    'user_id' =>$val_ui,
                                ));
                            }
                        }
                    }
                    $return['state'] =1;
                    $return['message'] = array(Ec::lang('操作成功'));
                    $db_apa->commit();
                }catch(Exception $e) {
                    $db_apa ->rollBack();
                }
            }
            die(Zend_Json::encode($return));
        }
    }

    
}