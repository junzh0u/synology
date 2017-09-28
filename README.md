    $ syno -host superman.synology.me:5000 -username admin -password mysecretpassword DSGetConfig | json_pp                                                
    {
       "bt_max_download" : 1024,
       "emule_enabled" : false,
       "emule_max_upload" : 20,
       "http_max_download" : 0,
       "bt_max_upload" : 32,
       "nzb_max_download" : 0,
       "default_destination" : "shared/Download",
       "emule_default_destination" : null,
       "unzip_service_enabled" : false,
       "ftp_max_download" : 0,
       "xunlei_enabled" : false,
       "emule_max_download" : 0
    }
    $ syno -host superman.synology.me:5000 -username admin -password mysecretpassword DSSetConfig -bt_max_download=2048
    $ syno -host superman.synology.me:5000 -username admin -password mysecretpassword DSGetConfig | json_pp                                                
    {
       "emule_max_upload" : 20,
       "bt_max_download" : 2048,
       "http_max_download" : 0,
       "unzip_service_enabled" : false,
       "default_destination" : "shared/Download",
       "ftp_max_download" : 0,
       "nzb_max_download" : 0,
       "emule_enabled" : false,
       "bt_max_upload" : 32,
       "xunlei_enabled" : false,
       "emule_default_destination" : null,
       "emule_max_download" : 0
    }
