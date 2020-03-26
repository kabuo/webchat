new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#appChatHist',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        // チャット履歴
        chatHist: [],
        
        // ユーザID(自分)
        myUserId: '',
        // ユーザID(相手)
        opUserId: '',

        // 送信メッセージ
        message: '',

        // 履歴最終日時(yyyyMMddHHmmssfff)
        lastHistTime: 0,

        // true：非活性・false：活性
        isDisabled: true,
    },

    // 算出プロパティ
    computed: {
        // 入力チェック
        validate() {
            // 未入力がある場合は、ボタン押下不可
            this.isDisabled = true
            if (this.message.length > 0)
                this.isDisabled = false
            return this.isDisabled
        }
    },

    // インスタンス作成時の処理
    created: function() {
        this.initInputValue();

        this.opUserId = 'Ｂ'
        //this.doFindAllChatHist();
    },

    // メソッド定義
    methods: {
        // チャット情報を全て取得する
        doFindAllChatHist : function(e) {
            axios.get('/findAllChatHist', {
                params: {
                    myUserId: this.myUserId,
                    opUserId: this.opUserId,
                }
            })
            .then(response => {
                if (response.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // サーバから取得したチャット履歴情報を設定
                    this.chatHist = response.data
                    
                    // TODO 
                    // my_flg
                    // user_id
                    // user_name
                    // send_date
                    // msg_str
                    // msg_stmp
                    // chk_read
                    
                }
            }).catch(error =>{
                const {
                    status, 
                    statusText,
                } = error.response;
                console.log(`Error! HTTP Status: ${status} ${statusText}`);
            })
        },

        // 新規チャット情報を登録する
        doAddChatHist : function(e) {
            //ログインユーザ登録要求API：/addLoginUser
            const params = new URLSearchParams();
            params.append('userId', this.userId)
            params.append('password', this.password)
            params.append('userName', this.userName)

            // 要求
            axios.post('/addChatHist', params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // チャット画面へ遷移
                    location.href = '//localhost:8080/chat';
                }
            }).catch(error =>{
                const {
                    status, 
                    statusText,
                } = error.response;
                console.log(`Error! HTTP Status: ${status} ${statusText}`);
            })
        },
        // 入力値を初期化する
        initInputValue() {
            myUserId = '',
            opUserId = '',
            message = '',
            this.isDisabled = true
        }
    }
})