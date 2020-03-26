new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#appLogin',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        loginUser: {
            // ユーザID
            userId: '',
            // パスワード
            password: '',
        },
        // ユーザID
        userId: '',
        // パスワード
        password: '',
        
        // true：非活性・false：活性
        isDisabled: true,
        
        // 入力エラータイプ
        errType: '',
    },

    // 算出プロパティ
    computed: {
        // 入力チェック
        validate() {
            this.isDisabled = true
            if (this.userId.length > 0 && this.password.length > 0)
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
        // ログインユーザ情報を取得する
        doFindLoginUser : function(e) {
            //ログインユーザ個別取得要求API：/findLoginUser
            axios.get('/findLoginUser', {
                params: {
                    userId: this.userId
                }
            })
            .then(response => {
                if (response.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // サーバから取得したログインユーザ情報をdataに設定する
                    this.loginUser = response.data
                    
                    this.errType = '';
                    // ユーザ未存在
                    if (this.loginUser.userId === '') {
                        this.errType = 'E01'
                    }
                    // パスワード異なる
                    else if (this.loginUser.password !== this.password) {
                        this.errType = 'E02'
                    }
                    else {
                        // チャット画面へ遷移
                        location.href = '//localhost:8080/webchat/chat.html';
                    }
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
            this.loginUser.userId = ''
            this.loginUser.password = ''
            this.userId = ''
            this.password = ''
            this.isEntered = false
            errType = ''
        }
    }
})