package models

import (
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Pedidos struct {
	NomeBeneficiario           string `json:"nomeBeneficiario"`
	Cpf                        string `json:"cpf"`
	NumeroGuiaPrestador        string `json:"numeroGuiaPrestador"`
	CodigoAutorizacao          int    `json:"codigoAutorizacao"`
	SenhaAutorizacao           string `json:"senhaAutorizacao"`
	DataAutorizacao            string `json:"dataAutorizacao"`
	DescricaoGlosa             string `json:"descricaoGlosa"`
	CodigoGlosa                string `json:"codigoGlosa"`
	DataExpiracaoAutorizacao   string `json:"dataExpiracaoAutorizacao"`
	TipoGuia                   string `json:"tipoGuia"`
	NumeroConselhoProfissional string `json:"numeroConselhoProfissional,int"`
	NomeProfissional           string `json:"nomeProfissional"`
	NumeroCarteira             string `json:"numeroCarteira"`
	AcomodacaoPlano            string `json:"acomodacaoPlano"`
	AtendimentoRN              string `json:"atendimentoRN"`
	CaraterAtendimento         string `json:"caraterAtendimento"`
	Cbos					   string `json:"cbos"`
	Cid                        string `json:"cid"`
	Cnes                       string `json:"cnes"`
	CodigoPrestadorNaOperadora string `json:"codigoPrestadorNaOperadora"`
	CodigoSolicitante          string `json:"codigoSolicitante"`
	ConselhoProfissional       string `json:"conselhoProfissional"`
	DataNascimento             string `json:"dataNascimento"`
	DataSolicitacao            string `json:"dataSolicitacao"`
	IndicacaoAcidente          string `json:"indicacaoAcidente"`
	IndicacaoClinica		   string `json:"indicacaoClinica"`
	MotivoEncerramentoAtendimento string `json:"motivoEncerramentoAtendimento"`
	NomeContratado             string `json:"nomeContratado"`
	NomeOperadora              string `json:"nomeOperadora"`
	NomeSolicitante			   string `json:"nomeSolicitante"`
	NumeroGuiaPrincipal        string `json:"numeroGuiaPrincipal"`
	Observacao                 string `json:"observacao"`
	Plano                      string `json:"plano"`
	RegistroANS				   string `json:"registroANS"`
	TipoAtendimento            string `json:"tipoAtendimento"`
	TipoConsulta	           string `json:"tipoConsulta"`
	TipoTransacao			   string `json:"tipoTransacao"`
	Uf						   string `json:"uf"`
	ValidadeCarteirinha        string `json:"validade_carteirinha"`
	ValidadeSenha			   string `json:"validadeSenha"`
}

//Função que verifica se já existe um paciente cadastrado, caso contrário, realiza a inserção no Clinux
func CreatePedidos(pedidos *Pedidos, p []byte) int {
	var cdPedido int

	db := config.ConnectDB()

	if len(pedidos.NumeroGuiaPrestador) > 0 {
		//Inserir pedido
		err := db.QueryRow(`insert into pedidos (ds_beneficiario, ds_cpf, ds_guia_prestador,
			 cd_autorizacao, ds_senha_autorizacao, dt_autorizacao, ds_glosa, cd_glosa,
			 dt_expiracao_autorizacao, ds_tipo_guia, ds_crm, ds_medico, nr_carteira, ds_guia_principal, 
             dt_validade_carteira, cd_prestador_na_operadora, ds_solicitante, cd_solicitante,
			 ds_conselho, ds_cbos, ds_acomodacao_plano, ds_atendimento_rn, ds_carater_atendimento, 
			 ds_cid, ds_cnes, dt_nascimento, ds_indicacao_acidente, ds_indicacao_clinica, 
			 ds_motivo_encerramento_atendimento, ds_observacao, ds_plano, ds_registro_ans, 
			 ds_tipo_atendimento, ds_tipo_consulta, ds_tipo_transacao, ds_uf, dt_solicitacao_guia, 
			 dt_validade_senha, json_envio) values
			  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
				$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39) returning id`,
			pedidos.NomeBeneficiario, pedidos.Cpf, pedidos.NumeroGuiaPrestador, pedidos.CodigoAutorizacao, pedidos.SenhaAutorizacao,
			pedidos.DataAutorizacao, pedidos.DescricaoGlosa, pedidos.CodigoGlosa, pedidos.DataExpiracaoAutorizacao,
			pedidos.TipoGuia, pedidos.NumeroConselhoProfissional, pedidos.NomeProfissional, pedidos.NumeroCarteira, pedidos.NumeroGuiaPrincipal,
			pedidos.ValidadeCarteirinha, pedidos.CodigoPrestadorNaOperadora, pedidos.NomeSolicitante, pedidos.CodigoPrestadorNaOperadora,
			pedidos.ConselhoProfissional, pedidos.Cbos, pedidos.AcomodacaoPlano, pedidos.AtendimentoRN, pedidos.CaraterAtendimento, pedidos.Cid,
			pedidos.Cnes, pedidos.DataNascimento, pedidos.IndicacaoAcidente, pedidos.IndicacaoClinica,pedidos.MotivoEncerramentoAtendimento, 
			pedidos.Observacao, pedidos.Plano, pedidos.RegistroANS, pedidos.TipoAtendimento, pedidos.TipoConsulta,
			pedidos.TipoTransacao, pedidos.Uf, pedidos.DataSolicitacao,pedidos.ValidadeSenha,p).Scan(&cdPedido)
		if err != nil {
			//handle error
			fmt.Printf("Error during insert pedidos: %v", err)
			//log.Fatalf("Error during insert pedidos: %v", err)
		}
	}

	defer db.Close()

	return cdPedido
}
