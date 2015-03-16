### Ad Tracker API

[![Build Status](https://api.travis-ci.org/wuiscmc/ads.svg)](https://travis-ci.org/wuiscmc/ads)

### Retrieve an Ad in VAST2 format
```
GET https://ads-api-test.herokuapp.com/v1/ads/1
```
##### Response
```xml
<VAST version="2.0">
    <Ad id="3" >
        <InLine>
            <AdSystem version="2.0.0">
                <![CDATA[ AdService deluxe ]]>
            </AdSystem>
            <AdTitle>
              <![CDATA[ title3 ]]>
            </AdTitle>
            <Impression id="AT">
                <![CDATA[ https://ads-api-test.herokuapp.com/v1/track?uid=3&action=videoClientSideImpression ]]>
            </Impression>
            <Creatives>
                <Creative id="3" sequence="1">
                    <Linear>
                        <Duration/>
                        <MediaFiles>

                            <MediaFile width="720" height="405" delivery="progressive" type="video/mp4">
                                <![CDATA[ http://ad-videos.host.com/videos/video5.mp4 ]]>
                            </MediaFile>

                            <MediaFile width="720" height="405" delivery="progressive" type="video/mp4">
                                <![CDATA[ http://ad-videos.host.com/videos/video6.mp4 ]]>
                            </MediaFile>

                        </MediaFiles>
                        <TrackingEvents>
                            <Tracking event="start">
                                <![CDATA[ https://ads-api-test.herokuapp.com/v1/track?uid=3&action=start ]]>
                            </Tracking>
                            <Tracking event="complete">
                                <![CDATA[ https://ads-api-test.herokuapp.com/v1/track?uid=3&action=complete ]]>
                            </Tracking>
                        </TrackingEvents>
                        <VideoClicks>
                            <ClickThrough>
                                <![CDATA[ https://ads-api-test.herokuapp.com/v1/clk?uid=3&action=videoClickThrough ]]>
                            </ClickThrough>
                        </VideoClicks>
                    </Linear>
                </Creative>
            </Creatives>
        </InLine>
    </Ad>
</VAST>
```

### Tracking events
```
POST https://ads-api-test.herokuapp.com/track?uid=3&action={complete|start|videoClientSideImpression}
```
#### Response
```
HTTP/1.1 204 No Content
```

