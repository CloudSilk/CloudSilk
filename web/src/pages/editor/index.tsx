import Editor from '@/form/components/common/AtaliEditor/editor';
import { Card, Col, Row } from 'antd';
import { useState } from 'react';

export default () => {
    //{ "time": 1681193234430, "blocks": [{ "id": "1rzorKy4Z5", "type": "video", "data": { "file": { "url": "/api/core/file/download?id=9cf8676e-8f9c-49d1-adb7-5ea333b2ea02" }, "caption": "", "withBorder": false, "stretched": false, "withBackground": false } }, { "id": "RL_B1J478I", "type": "image", "data": { "file": { "url": "/api/core/file/download?id=4ebd4e7c-5df0-4e3d-bae6-328a4b368109" }, "caption": "", "withBorder": false, "stretched": false, "withBackground": false } }, { "id": "gbqH-bX-ey", "type": "video", "data": { "file": { "url": "/api/core/file/download?id=552033e8-e7a6-44a5-8d7c-0f27adcc6f9e" }, "caption": "", "withBorder": false, "stretched": false, "withBackground": false } }], "version": "2.26.5" }
    const data = { "time": 1681206399449, "blocks": [{ "id": "FpGP9B-VtZ", "type": "header", "data": { "text": "<span style=\"vertical-align: top; font-family: 宋体, SimSun; margin-top: 14px; margin-bottom: 14px; line-height: 1.4em;\"><font>庄严在自治区党委党校(行政学院)调研座谈时强调</font></span>", "level": 1 }, "tunes": { "anyTuneName": { "alignment": "center" } } }, { "id": "NdH23HBcP2", "type": "header", "data": { "text": "<font>深入学习贯彻习近平总书记重要讲话精神 奋力推动党校工作高质量发展</font>", "level": 3 }, "tunes": { "anyTuneName": { "alignment": "center" } } }, { "id": "Uk_gMVI9kq", "type": "paragraph", "data": { "text": "<span style=\"color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: justify; text-indent: 38px;\">&nbsp; &nbsp; 4月7日,自治区党委常务副书记、自治区党委党校(行政学院)校(院)长庄严来到自治区党委党校(行政学院)调研并召开座谈会。他强调,要深入学习贯彻习近平总书记在中央党校建校90周年庆祝大会暨2023年春季学期开学典礼上的重要讲话精神,学习贯彻全国党校(行政学院)校(院)长会议精神,贯彻落实自治区党委和王君正书记的部署要求,始终坚守党校初心,把握办学治校规律,奋力推动全区党校工作高质量发展。</span>" }, "tunes": { "anyTuneName": { "alignment": "left" } } }, { "id": "GcUSBo-_S-", "type": "paragraph", "data": { "text": "<span style=\"color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: justify; text-indent: 38px;\">&nbsp; &nbsp; 庄严先后深入中青年班课堂、党史党建教研部、科学社会主义教研部,饶有兴致地观摩教学,与大家交流实践体会。他勉励学员要坚持不懈用习近平新时代中国特色社会主义思想武装头脑、指导实践、推动工作,弘扬理论联系实际的学风,不断提升理论素养和能力本领,在全面建设社会主义现代化新西藏的火热实践中建功立业。要求党校把立德树人作为中心工作,坚持价值塑造、能力培养、知识传授“三位一体”,努力建设在全国有特色、民族地区有地位、涉藏地区创一流的西藏党校。</span>" }, "tunes": { "anyTuneName": { "alignment": "left" } } }, { "id": "l_I-Apv2A7", "type": "paragraph", "data": { "text": "<span style=\"color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: justify; text-indent: 38px;\">&nbsp; &nbsp; 在座谈会上,庄严强调,全区党校系统要坚持把学习贯彻习近平总书记重要讲话精神作为重大政治任务,认真组织开展形式多样的大学习、大讨论,深刻领会把握为党育才、为党献策的党校初心,深刻领会把握坚守党校初心的丰富内涵,深刻领会把握坚持党对党校工作全面领导的根本要求,做到学思用贯通、知信行统一。</span>" }, "tunes": { "anyTuneName": { "alignment": "left" } } }, { "id": "EJGVjB0fTL", "type": "paragraph", "data": { "text": "<span style=\"color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: justify; text-indent: 38px;\">&nbsp; &nbsp; (下转第三版) (紧接第一版) 要认真对标对表习近平总书记重要讲话提出的目标要求,全面梳理党校工作中的短板弱项,制定工作方案,压实责任分工,做到目标、任务、措施、责任、时限具体化,在补短板、强弱项、扬优势上下功夫、见实效。要始终牢记“国之大者”,自觉服务于西藏长治久安和高质量发展,进一步加强马克思主义理论教育培训,更加扎实做好习近平新时代中国特色社会主义思想宣传研究阐释工作,广泛深入宣传党的主张,持续深化质量立校工作。要全面推进从严治校,严守政治纪律和政治规矩,严以治校、严以治教、严以治学,大力营造朴素清朗的学风、忠诚敬业的教风、风清气正的校风,以实际行动坚决捍卫“两个确立”、做到“两个维护”。</span><span style=\"vertical-align: top; color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: justify; text-indent: 38px;\">(记者 王雅慧 实习生 王靖宇)</span>" }, "tunes": { "anyTuneName": { "alignment": "left" } } }, { "id": "SFjTQhilKN", "type": "paragraph", "data": { "text": "<span style=\"color: rgb(0, 0, 0); font-family: 宋体, SimSun; font-size: 19px; text-align: right; text-indent: 38px;\">(编辑:谭露婷)</span>" }, "tunes": { "anyTuneName": { "alignment": "right" } } }], "version": "2.26.5" }
    const [result, setResult] = useState<any>(data);
    return (
        <Row>
            <Col span={16}>
                <Editor preview={true} value={JSON.stringify(result)} onChange={(val) => {
                    setResult(JSON.parse(val));
                }} /></Col>
            <Col span={8}>
                <Card style={{ width: '90%', marginLeft: 'auto', marginRight: 'auto', marginTop: 10 }}><code lang="json">{JSON.stringify(result)}</code></Card>
            </Col>
        </Row>

    );
};