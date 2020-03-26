new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#appSignup',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        // ログインユーザ情報
        loginUsers: [],
        // ユーザID
        userId: '',
        // パスワード
        password: '',
        // ユーザ名
        userName: '',
        // ユーザ名
        passwordConfirm: '',
        
        // true：非活性・false：活性
        isDisabled: true,

        // 入力エラータイプ
        errType: '',
    },

    // 算出プロパティ
    computed: {
        // 入力チェック
        validate() {
            // 未入力がある場合は、ボタン押下不可
            this.isDisabled = true
            if (this.userId.length > 0 && this.password.length > 0 && this.userName.length > 0 && this.passwordConfirm.length > 0 )
                this.isDisabled = false
            return this.isDisabled
        }
    },

    // インスタンス作成時の処理
    created: function() {
        this.initInputValue();
    },

    // メソッド定義
    methods: {
        // ログインユーザ情報を登録する
        doAddLoginUser : function(e) {
           this.errType = '';

           // パスワード確認不一致
           if (this.password != this.passwordConfirm) {
               this.errType = 'E01' 
           }
           else {
               // ユーザID重複チェック
               axios.get('/findAllLoginUsers')
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } 
                    else {
                        // サーバから取得したログインユーザ情報をdataに設定する
                        this.loginUsers = response.data
                        var index = this.loginUsers.findIndex(x => x.userId == this.userId);
                        if (index !== -1) {
                            this.errType = 'E02'
                        }
                        else {
                            //ログインユーザ登録要求API：/addLoginUser
                            const params = new URLSearchParams();
                            params.append('userId', this.userId)
                            params.append('password', this.password)
                            params.append('userName', this.userName)

                            // 要求
                            axios.post('/addLoginUser', params)
                            .then(response => {
                                if (response.status != 200) {
                                    throw new Error('レスポンスエラー')
                                } else {
                                    // チャット画面へ遷移
                                    location.href = '//localhost:8080/webchat/chat.html';
                                }
                            }).catch(error =>{
                                const {
                                    status, 
                                    statusText,
                                } = error.response;
                                console.log(`Error! HTTP Status: ${status} ${statusText}`);
                            })
                        }
                    }
                }).catch(error =>{
                    const {
                        status, 
                        statusText,
                    } = error.response;
                    console.log(`Error! HTTP Status: ${status} ${statusText}`);
                })
           }
        },
        // 入力値を初期化する
        initInputValue() {
            this.userId = ''
            this.password = ''
            this.userName = ''
            this.passwordConfirm = ''
            this.isDisabled = true
            this.errType = ''
        }
    }
})